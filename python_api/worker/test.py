from celery import Celery, Task


app = Celery("test", broker="amqp://guest@localhost//")


@app.task(base=Task, bind=True, serializer="json")
def run_celery_test(self):
    id: str = self.request.id

    print(id)

    return "hello world"
