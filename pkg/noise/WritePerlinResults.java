import java.io.FileWriter;   // Import the FileWriter class
import java.io.IOException;  // Import the IOException class to handle errors
import java.lang.Math;

public class WritePerlinResults {
    public static void main(String[] args) {
        try {
            FileWriter myWriter = new FileWriter("perlinResults.txt");
            for (int i = 0; i < 1000; i++){
                double x = Math.random();
                double y = Math.random();
                double z = Math.random();
                double noise = Perlin.noise(x, y, z);
                myWriter.write(String.valueOf(x) + "," + String.valueOf(y) + "," + String.valueOf(z) + "," + String.valueOf(noise) + "\n");
            }
            myWriter.close();
            System.out.println("Successfully wrote to the file.");
        } catch (IOException e) {
            System.out.println("An error occurred.");
            e.printStackTrace();
        }
    }
}

