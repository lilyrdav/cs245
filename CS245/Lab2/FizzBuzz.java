public class FizzBuzz {
    public static void main (String[] args) {
        int count = 0;
        for (int i = 0; i < 100; i++) {
            count++;
            if (((count%3) == 0) && ((count%5) == 0)){
                System.out.println(count + " FizzBuzz");
            }
            else if ((count%5) == 0) {
                System.out.println(count + " Buzz");
            }
            else if ((count%3) == 0) {
                System.out.println(count + " Fizz");
            }
            else {
                System.out.println(count);
            }
        }
    }
}