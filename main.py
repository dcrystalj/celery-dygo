from celery import Celery
from kombu import Queue
import time

app = Celery("tasks", broker="redis://localhost:6379/0")
app.conf.task_queues = (
    Queue(
        name="celery",
    ),
)
app.conf.task_create_missing_queues = False


def main():
    app.send_task("myproject.apps.myapp.tasks.mytask", args=[3, 2], queue="celery")
    print("enqued")
    time.sleep(1)
    app.send_task("myproject.apps.myapp.tasks.mytask5", args=[1, 1], queue="celery")
    print("enqued 5")


if __name__ == "__main__":
    main()
