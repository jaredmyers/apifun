from fastapi import APIRouter, Depends
from storage.initdb import get_session, get_async_session
from storage.models import User, Store, Food
from models.models import CreateUserRequest, DeleteUserRequest
from sqlmodel import select
from fastapi.responses import JSONResponse
from sqlalchemy.orm import selectinload, joinedload


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
async def create_user(n: CreateUserRequest, session=Depends(get_async_session)):

    # if username already exists, return error
    statement = select(User).where(User.username == n.username)
    user_check = await session.execute(statement)
    user_check = user_check.scalars().first()

    if user_check is not None:
        return JSONResponse(status_code=422, content={"msg": "username already exists"})

    # Create username
    new_user = User(username=n.username, pw_hash=n.pw_hash)
    session.add(new_user)
    await session.commit()
    await session.refresh(new_user)
    return JSONResponse(status_code=200, content={"user_id": new_user.id})


@router.delete("/users")
async def delete_user(n: DeleteUserRequest, session=Depends(get_async_session)):
    statement = select(User).filter(User.username == n.username)
    results = await session.execute(statement)
    user = results.scalars().first()

    if user is None:
        return JSONResponse(status_code=422, content={"msg": f"{n.username} does not exist"})

    await session.delete(user)
    await session.commit()
    return JSONResponse(status_code=200, content={"deleted": user.username})


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
