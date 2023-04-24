from main import app, base
import pytest
from fastapi.testclient import TestClient
from sqlmodel import Session, SQLModel, create_engine
from httpx import AsyncClient


#client = TestClient(app)


@pytest.mark.asyncio
async def test_create_user():

    async with AsyncClient(app=app, base_url="http://localhost:8000"+base) as ac:
        response = await ac.post("/users", json={"username": "tester", "pw_hash": "pwman"})

    assert response.status_code == 200

    async with AsyncClient(app=app, base_url="http://localhost:8000"+base) as ac:
        response = await ac.post("/users", json={"username": "tester", "pw_hash": "pwman"})

    assert response.status_code == 422


@pytest.mark.asyncio
async def test_get_users():
    async with AsyncClient(app=app, base_url="http://localhost:8000"+base) as ac:
        response = await ac.get("/users")

    assert response.status_code == 200
