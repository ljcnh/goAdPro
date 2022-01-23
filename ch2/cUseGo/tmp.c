#include "tmp.h"
#include <stdio.h>

int foo() {
    return add(1, 1);
}

int main() {
    printf("%lld\n",foo());
    return 0;
}