import java.text.*;
import java.util.*;
public class JavaSpeed 
{
    public static void main(String args[]){
    int limit=1000;
    long mylist[]=new long[limit];
    long start = System.currentTimeMillis();
    for(int i=0;i<limit;i++){
            mylist[i]=i;
            System.out.println(mylist[i]);
    }

    long duration = System.currentTimeMillis()-start;
    NumberFormat formatter = new DecimalFormat("#0.00000");
System.out.print("Execution time is =" + formatter.format(duration) + " ms");
    }
}
