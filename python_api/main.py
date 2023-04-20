from fastapi import FastAPI
from api import users

app = FastAPI(
    title="apifun",
    description="morefun",
    )

app.include_router(users.router)
