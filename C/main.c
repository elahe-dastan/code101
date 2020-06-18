#include <stdio.h>
#include <stdlib.h>

char *change(char input[50]) {
    char *output;
    int begin, end, count = 0;

    //Calculating string length
    while (input[count] != '\0') {
        count++;
    }

    output= (char*)malloc((count + 2));

    end = count - 1;

    output[0] = 'A';

    for (begin = 1; begin <= count; begin++) {
        output[begin] = input[end];
        end--;
    }

    output[begin] = 'a';
    output[begin+1] = '\0';

    return (char*)output;
}

int main() {
    printf("%s", change("hello"));
    return 0;
}
