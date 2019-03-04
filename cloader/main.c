#include <stdlib.h>
#include <stdio.h>
#include <dlfcn.h>

int main() {
    void *handle;
    void (*doit)();

    handle = dlopen("out/cloader.so", RTLD_LAZY);
    if (!handle) {
        printf("error opening SO\n");
        fputs(dlerror(), stderr);
        exit(1);
    }
    doit = dlsym(handle, "doit");
    doit();
    dlclose(handle);
    return 0;
}
