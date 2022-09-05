#include <stdio.h>
#include <time.h>       // for clock_t, clock(), CLOCKS_PER_SEC
#include <unistd.h>     // for sleep()
 
// main function to find the execution time of a C program
int main()
{
    // to store the execution time of code
    double time_spent = 0.0;
 
    clock_t begin = clock();
 
    // do some stuff here
    int limit=1000;
    int list [limit];
    for (int i=0;i<limit;i++){
        list[i]=i;
        printf("%d\n",list[i]);
    }
 
    clock_t end = clock();
 
    // calculate elapsed time by finding difference (end - begin) and
    // dividing the difference by CLOCKS_PER_SEC to convert to seconds
    time_spent += (double)(end - begin);
 
    printf("Execution Time is =%f ms", time_spent);
 
    return 0;
}