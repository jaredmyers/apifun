from fastapi import APIRouter, Depends
from storage.initdb import get_session
from storage.models import User
from sqlmodel import select
from fastapi.responses import JSONResponse


router = APIRouter()


# == User Routes for Administration ==

# gets all users
@router.get("/users")
async def get_users(session=Depends(get_session)):
    statement = select(User)
    result = session.execute(statement).all()
    return result


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


# get user by username

# create list of users
