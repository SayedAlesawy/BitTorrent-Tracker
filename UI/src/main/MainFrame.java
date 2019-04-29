package main;

import javafx.application.Platform;
import javafx.geometry.Point2D;
import javafx.scene.canvas.Canvas;
import javafx.scene.canvas.GraphicsContext;
import javafx.scene.control.Tooltip;
import javafx.scene.input.KeyCode;
import javafx.scene.layout.Pane;
import javafx.scene.paint.Color;
import javafx.scene.paint.Paint;
import javafx.scene.text.FontSmoothingType;
import javafx.scene.transform.Affine;
import javafx.scene.transform.NonInvertibleTransformException;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.Response;
import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;

import java.io.IOException;
import java.util.ArrayList;

class MainFrame extends Pane {
    private static final String INFO_HASH = "InfoHash";
    private static final String PEERS = "PeerInfo";
    private static final String PEER = "Peer";
    private static final String STAT = "Stat";
    private static final String ID = "ID";
    private static final String IP = "IP";
    private static final String PORT = "Port";
    private static final String UPLOADED = "Uploaded";
    private static final String DOWNLOADED = "Downloaded";
    private static final String LEFT = "Left";
    private static final String EVENT = "Event";
    private static final String STARTED_EVENT = "started";
    private static final String STOPPED_EVENT = "stopped";
    private static final String COMPLETED_EVENT = "completed";

    private static final double PEER_RADIUS = 20;
    private static final long SLEEP_DURATION = 1000;

    boolean zoomIn = false;
    boolean zoomOut = false;

    ArrayList<Swarm> oldData;
    ArrayList<Swarm> data;
    volatile ArrayList<Swarm> newData;
    volatile boolean isConnected;
    volatile String link;
    volatile Affine transform;
    Sidebar sidebar;
    Canvas canvas;

    private double initX;
    private double initY;
    private double oldX;
    private double oldY;

    MainFrame() {
        newData = new ArrayList<>();
        sidebar = new Sidebar();

        isConnected = false;
        Thread t = new Thread(() -> connect());
        t.setDaemon(true);
        t.start();
        sidebar.setHandler((e) -> {
            if (isConnected) {
                isConnected = false;
                sidebar.setConnected(false);
            } else {
                link = sidebar.getLink();
                isConnected = true;
                sidebar.setConnected(true);
            }
        });
        transform = new Affine();
        canvas = new Canvas();

        heightProperty().addListener((observable, oldVal, newVal) -> {
            canvas.setHeight(newVal.doubleValue());
            transform.setTy(newVal.doubleValue() / 2);
            paint();
        });
        widthProperty().addListener((observable, oldVal, newVal) -> {
            canvas.setWidth(newVal.doubleValue());
            transform.setTx(newVal.doubleValue() / 2);
            paint();
        });
        getChildren().add(canvas);
        getChildren().add(sidebar);

        setOnKeyPressed((e) -> {
            if (e.getCode() == KeyCode.I) {
                zoomIn = true;
            } else if (e.getCode() == KeyCode.O) {
                zoomOut = true;
            }
        });
        setOnKeyReleased((e) -> {
            if (e.getCode() == KeyCode.I) {
                zoomIn = false;
            } else if (e.getCode() == KeyCode.O) {
                zoomOut = false;
            }
        });

        Thread zoomThread = new Thread(() -> {
            while (true) {
                if (zoomIn) {
                    transform.setMyy(transform.getMyy() * 1.01);
                    transform.setMxx(transform.getMxx() * 1.01);
                    transform.setTx(transform.getTx());
                    transform.setTy(transform.getTy());
                } else if (zoomOut) {
                    transform.setMyy(transform.getMyy() * 0.99);
                    transform.setMxx(transform.getMxx() * 0.99);
                    transform.setTx(transform.getTx());
                    transform.setTy(transform.getTy());
                }
                if (zoomIn || zoomOut) {
                    Platform.runLater(() -> paint());
                    try {
                        Thread.sleep(10);
                    } catch (InterruptedException e) {
                        e.printStackTrace();
                    }
                } else {
                    try {
                        Thread.sleep(100);
                    } catch (InterruptedException e) {
                        e.printStackTrace();
                    }
                }
            }
        });
        zoomThread.setDaemon(true);
        zoomThread.start();

        setOnMouseClicked((e) -> {
            canvas.requestFocus();
        });

        setOnMousePressed((e) -> {
            initX = e.getSceneX();
            initY = e.getSceneY();
            oldX = transform.getTx();
            oldY = transform.getTy();
        });

        setOnMouseDragged((e) -> {
            transform.setTx(oldX + e.getSceneX() - initX);
            transform.setTy(oldY + e.getSceneY() - initY);
            paint();
        });
        Tooltip tp = new Tooltip();
        setOnMouseMoved((e) -> {
            double x = e.getSceneX();
            double y = e.getSceneY();
            boolean flag = false;
            for (Swarm s : oldData) {
                for (Peer p : s.peers) {
                    double px = p.point.getX();
                    double py = p.point.getY();
                    if (x >= px && y >= py && x <= px + PEER_RADIUS && y <= py + PEER_RADIUS) {
                        flag = true;
                        tp.setText("ID: " + p.id + "\nIP: " + p.ip + ":" + p.port +
                                "\n" + "Status: " + p.event + "\n" +
                                "Uploaded: " + p.uploaded + "\n" +
                                "Downloaded: " + p.downloaded + "\n" +
                                "Left: " + p.left);
                        tp.show(this, e.getScreenX(), e.getScreenY());
                    }
                }
                if (!flag) tp.hide();
            }
        });
    }

    private void paint() {
        data = new ArrayList<>(newData);
        GraphicsContext gc = canvas.getGraphicsContext2D();
        gc.setFontSmoothingType(FontSmoothingType.GRAY);
        gc.setTransform(1, 0, 0, 1, 0, 0);
        gc.clearRect(0, 0, canvas.getWidth(), canvas.getHeight());
        gc.setTransform(transform);

        double dInc = PEER_RADIUS * 20;
        double d = 0;
        double angle = 0;
        int k = 4;
        int j = 1;
        for (Swarm s : data) {
            paintSwarm(gc, s, d * Math.cos(angle), d * Math.sin(angle));
            angle += 2 * Math.PI / k;
            j--;
            if (j == 0) {
                d += dInc;
                angle += Math.PI / k;
                k *= 2;
                j = k;
            }
        }
        oldData = new ArrayList<>(data);
    }

    void paintSwarm(GraphicsContext gc, Swarm s, double x, double y) {
        double dInc = PEER_RADIUS * 1.5;
        double d = PEER_RADIUS * 1.5;
        double angle = 0;

        double minX = Float.POSITIVE_INFINITY;
        double minY = Float.POSITIVE_INFINITY;
        double maxX = Float.NEGATIVE_INFINITY;
        double maxY = Float.NEGATIVE_INFINITY;

        int k = 4;
        int j = k;
        for (Peer p : s.peers) {
            double pX = x + d * Math.cos(angle);
            double pY = y + d * Math.sin(angle);
            paintPeer(gc, p, pX, pY);
            angle += 2 * Math.PI / k;
            j--;
            if (j == 0) {
                d += dInc;
                angle += Math.PI / k;
                k *= 2;
                j = k;
            }

            if (pX < minX) minX = pX;
            if (pY < minY) minY = pY;
            if (pX + PEER_RADIUS > maxX) maxX = pX + PEER_RADIUS;
            if (pY + PEER_RADIUS > maxY) maxY = pY + PEER_RADIUS;
        }
        minX -= 10;
        maxX += 10;
        minY -= 10;
        maxY += 10;
        gc.strokeRect(minX, minY, maxX - minX, maxY - minY);
        gc.fillText(s.info, minX + 10, minY - 10);
    }

    void paintPeer(GraphicsContext gc, Peer p, double x, double y) {
        gc.strokeOval(x, y, PEER_RADIUS, PEER_RADIUS);
        Paint paint = gc.getFill();
        if (p.event.equals(STARTED_EVENT)) {
            gc.setFill(Color.BLUE);
        } else if (p.event.equals(STOPPED_EVENT)) {
            gc.setFill(Color.RED);
        } else if (p.event.equals(COMPLETED_EVENT)) {
            gc.setFill(Color.GREEN);
        } else {
            gc.setFill(Color.TRANSPARENT);
        }
        gc.fillOval(x, y, PEER_RADIUS, PEER_RADIUS);
        gc.setFill(paint);
        p.point = transform.transform(x, y);
    }

    void connect() {
        while (true) {
            if (isConnected) {
                OkHttpClient client = new OkHttpClient();
                Request request = new Request.Builder()
                        .url(link)
                        .get()
                        .addHeader("Content-Type", "application/json")
                        .addHeader("cache-control", "no-cache")
                        .addHeader("Postman-Token", "5a5b3dd2-48e8-4a87-9062-405a804d8d2d")
                        .build();

                Response response;
                try {
                    response = client.newCall(request).execute();
                    String content = response.body().string();
                    System.out.println(content);
                    if (content.startsWith("null")) {
                        newData.clear();
                    } else {
                        JSONArray jsonArray = new JSONArray(content);
                        newData = jsonToData(jsonArray);
                    }
                    Platform.runLater(() -> paint());
                } catch (IOException | JSONException e) {
                    Platform.runLater(() -> {
                        isConnected = false;
                        sidebar.setConnected(false);
                    });
                    e.printStackTrace();
                    return;
                }
                try {
                    Thread.sleep(SLEEP_DURATION);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
        }
    }

    private ArrayList<Swarm> jsonToData(JSONArray jsonArray) {
        ArrayList<Swarm> newData = new ArrayList<>();
        for (Object object : jsonArray) {
            JSONObject jsonObject = (JSONObject) object;
            Swarm swarm = new Swarm();
            swarm.peers = new ArrayList<>();
            swarm.info = jsonObject.getString(INFO_HASH);
            JSONArray peers = jsonObject.getJSONArray(PEERS);
            for (Object p : peers) {
                Peer peer = new Peer();
                JSONObject peerInfoJson = (JSONObject) p;
                JSONObject peerJson = peerInfoJson.getJSONObject(PEER);
                JSONObject statJson = peerInfoJson.getJSONObject(STAT);
                peer.id = peerJson.getInt(ID);
                peer.ip = peerJson.getString(IP);
                peer.port = peerJson.getInt(PORT);
                peer.uploaded = statJson.getInt(UPLOADED);
                peer.downloaded = statJson.getInt(DOWNLOADED);
                peer.left = statJson.getInt(LEFT);
                peer.event = statJson.getString(EVENT);
                swarm.peers.add(peer);
            }
            newData.add(swarm);
        }
        return newData;
    }
}
