from fastapi import APIRouter, Depends
from storage.initdb import get_session
from storage.schemas import User, Team
from sqlmodel import select
from fastapi.responses import JSONResponse

from storage.initdb import engine
from sqlmodel import Session, create_engine

router = APIRouter()


# == Users Endpoint ==
@router.get("/users")
async def get_users(session=Depends(get_session)):
    statement = select(User)
    result = session.execute(statement).all()
    return result


@router.post("/users")
async def create_user(new_user: User, session=Depends(get_session)):
    session.add(new_user)
    session.commit()
    return JSONResponse(status_code=200, content={"user_id": "test"})


@router.get("/user/{id}")
def get_user(user_id: int, session=Depends(get_session)):
    statement = select(User).filter(User.id == id)
    result = session.execute(statement).all()
    return result


@router.get("/testing")
def testing():

    with Session(engine) as session:
        team_preventers = Team(name="Preventers", headquarters="Sharp Tower")
        team_z_force = Team(name="z-force", headquarters="tower honai")
        session.add(team_preventers)
        session.add(team_z_force)
        session.commit()

        user_dp = User(fname="Deapond", lname="noth",
                       email="mail@mail.com",
                       team_id=team_z_force.id)

        user_r = User(fname="Rusty", lname="Ljjs",
                      email="mail@mail.com",
                      team_id=team_preventers.id)

        session.add(user_dp)
        session.add(user_r)
        session.commit()

        session.refresh(user_dp)
        session.refresh(user_r)

        print("Created:", user_dp)
        print("Created:", user_r)

        b = User(fname="nacho", lname="libre", email="lskjdf")

        return None


@router.get("/t")
def t():
    with Session(engine) as session:
        statement = select(User).where(User.fname == "larry")
        result = session.exec(statement).one()
        print(result)
        print(type(result))
        print(result.team)
        print(type(result.team))
        return {1: result, 2: result.team}


@router.get("/p")
def p(session=Depends(get_session)):
    statement = select(User).where(User.fname == "larry")
    result = session.execute(statement).one()
    print(result)
    print(type(result))
    print(result.team)
    return None

