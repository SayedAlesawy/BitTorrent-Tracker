package main;

import javafx.application.Application;
import javafx.scene.Scene;
import javafx.stage.Stage;

public class Main extends Application {
    @Override
    public void start(Stage primaryStage) throws Exception {
        try {
            MainFrame root = new MainFrame();
            Scene scene = new Scene(root, 1000, 600);
            root.minWidthProperty().bind(scene.widthProperty());
            root.minHeightProperty().bind(scene.heightProperty());
            primaryStage.setScene(scene);
            primaryStage.show();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    public static void main(String[] args) {
        launch(args);
    }
}
