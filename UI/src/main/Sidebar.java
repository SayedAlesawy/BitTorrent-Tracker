package main;

import javafx.event.ActionEvent;
import javafx.event.EventHandler;
import javafx.geometry.Insets;
import javafx.geometry.Pos;
import javafx.scene.control.Button;
import javafx.scene.control.TextField;
import javafx.scene.layout.HBox;
import javafx.scene.layout.VBox;

public class Sidebar extends VBox {

    private TextField ipTextField;
    private Button connectButton;

    public Sidebar() {
        // set up style
        setPadding(new Insets(10));
        setSpacing(10);

        HBox hbox1 = new HBox();
        hbox1.setAlignment(Pos.CENTER);
        hbox1.setSpacing(10);
        ipTextField = new TextField("http://localhost:3000/swarms");
        connectButton = new Button("Connect");
        connectButton.setDefaultButton(true);
        hbox1.getChildren().add(ipTextField);
        hbox1.getChildren().add(connectButton);
        getChildren().add(hbox1);
    }

    String getLink() {
        return ipTextField.getText();
    }

    void setConnected(boolean value) {
        if (value) connectButton.setText("Disconnect");
        else connectButton.setText("Connect");
    }

    void setHandler(EventHandler<ActionEvent> handler) {
        connectButton.setOnAction(handler);
    }
}
