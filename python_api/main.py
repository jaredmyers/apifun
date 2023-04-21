from fastapi import FastAPI
from api import users

app = FastAPI(
    title="apifun",
    description="apifun",
    )

base = "/api/v1"
app.include_router(users.router, prefix=base)
