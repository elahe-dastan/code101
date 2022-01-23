public class Main {
    public static void main(String[] args) {

        AsyncDatabaseWriter writer = new AsyncDatabaseWriter();
        Thread t1 =new Thread(writer);   // Using the constructor Thread(Runnable r)  
        t1.start();

        writer.ScheduleWrite("gobol");
    }
}

