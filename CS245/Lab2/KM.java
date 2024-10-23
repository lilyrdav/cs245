public class KM{
    public static void main (String[] args) {
        double number = Double.parseDouble(args[0]);
        String unit = args[1];
        double kilometers = number * 1.609;
        double miles = number / 1.609;
        if (unit.equals("m")) {
            System.out.println(kilometers + " km");
        }
        else if (unit.equals("k")) {
            System.out.println(miles + " miles");
        }
    }
}
