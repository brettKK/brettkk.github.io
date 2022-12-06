#include <pthread.h>
#include <stdio.h>

#define LOOP 1000000

pthread_mutex_t mutex;
int cnt = 0;
int cs1 = 0, cs2 = 0;

void *task(void *args) {
    while(1) {
        pthread_mutex_lock(&mutex);
        if (cnt >= LOOP) {
            pthread_mutex_unlock(&mutex);
            break;
        }
        cnt++;
        pthread_mutex_unlock(&mutex);
        if ((int)args == 1) cs1++;
        else cs2++;
    }
    return NULL;
}

int main() {
    pthread_mutex_init(&mutex, NULL);
    pthread_t tid1;
    pthread_t tid2;
    pthread_create(&tid1, NULL, task, (void*)1);
    pthread_create(&tid2, NULL, task, (void*)2);
    pthread_join(tid1, NULL);
    pthread_join(tid2, NULL);

    printf("cnt = %d, cs1 = %d, cs2 = %d,  total=%d\n", cnt, cs1, cs2, cs1 + cs2);
    return 0;
}