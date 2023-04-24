from fastapi import FastAPI
from api import users, foods, stores

app = FastAPI(
    title="food tracking",
    description="api",
    )

base = "/api/v1"
app.include_router(users.router, prefix=base)
app.include_router(foods.router, prefix=base)
app.include_router(stores.router, prefix=base)
