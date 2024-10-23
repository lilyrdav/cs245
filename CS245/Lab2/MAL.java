import java.util.ArrayList;

public class MAL extends ArrayList<Integer>{
    public static void main (String[] args) {
        MAL sal = new MAL();
        int number = 0;
        for (int i = 0; i < 20; i++) {
            number++;
            sal.add(number);
        }
        System.out.println(sal);
    }

    /**
     * This function takes two parameters, the start and the count. 
     * It returns a new instance of the class that contains the items
     * starting at the start index and with the given count. For example
     * if the current instance contains [0,1,2,3,4,5,6,7,8,9] and this 
     * method was called with the the parameters 1,3 it would return
     * a new instance of the class with [1,2,3]. With parameters
     * 4,5 it would return [4,5,6,7,8]
     * If the first number is less than 0, count backwards from the end.
     * For example -3,2 would give [7,8].
     * @param start the index (the first item is at index 0) 
     * of the first item in the current list to 
     * be copied into the new list. If start index is less than 0, 
     * start at length-start index.
     * @param count the number of items to be copied. If count is too large
     * return as many as are available.   If count is less than 0, return all items
     * to the end of the current list.
     * @return a new class containing a subset of the objects in the 
     * source class
     */
    public MAL slicer(int start, int count) {
        int[] array = new int[count];
        
        for (int i = 0; i < this.size(); i++) {
           if (this.get(start).equals() ) 
        }
    }
}