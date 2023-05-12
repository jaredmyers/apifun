from fastapi import APIRouter, Depends
from storage.initdb import get_session, get_async_session
from storage.models import User, Store, Food
from models.models import CreateUserRequest, DeleteUserRequest
from sqlmodel import select
from fastapi.responses import JSONResponse
from sqlalchemy.orm import selectinload, joinedload

from worker.test import run_celery_test
from fastapi.encoders import jsonable_encoder


router = APIRouter()


# == User Routes for Administration ==


# gets all users
@router.get("/users")
async def get_users(session=Depends(get_async_session)):
    statement = select(User)
    result = await session.execute(statement)
    return result.all()


@router.get("/users/{id}")
async def get_user_by_id(session=Depends(get_async_session)):
    statement = select(User).where(User.id == id)
    result = await session.execute(statement)
    return result.first()


@router.post("/users")
async def create_user(cU: CreateUserRequest, session=Depends(get_async_session)):

    # if username already exists, return error
    statement = select(User).where(User.username == cU.username)
    user_check = await session.execute(statement)
    user_check = user_check.scalars().first()

    if user_check is not None:
        return JSONResponse(status_code=422, content={"msg": "username already exists"})

    # Create username
    new_user = User(username=cU.username, pw_hash=cU.pw_hash)
    session.add(new_user)
    await session.commit()
    await session.refresh(new_user)
    return JSONResponse(status_code=200, content={"user_id": new_user.id})


@router.delete("/users")
async def delete_user(dU: DeleteUserRequest, session=Depends(get_async_session)):
    statement = select(User).filter(User.username == dU.username)
    results = await session.execute(statement)
    user = results.scalars().first()

    if user is None:
        return JSONResponse(status_code=422, content={"msg": f"{dU.username} does not exist"})

    await session.delete(user)
    await session.commit()
    return JSONResponse(status_code=200, content={"deleted": user.username})


@router.get("/celery/test")
async def test_celery():
    example = run_celery_test.delay()
    print(example)

    return JSONResponse(status_code=200, content={"celery": "yes"})


"""

# create new user
@router.post("/users")
async def create_user(new_user: User, session=Depends(get_session)):
    session.add(new_user)
    session.commit()
    return JSONResponse(status_code=200, content={"user_id": "test"})


# get user by id
@router.get("/user/{id}")
def get_user(user_id: int, session=Depends(get_session)):
    statement = select(User).filter(User.id == id)
    result = session.execute(statement).all()
    return result

"""

# get user by username

# create list of users
