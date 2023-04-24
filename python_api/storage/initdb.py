"""
from sqlmodel import create_engine, Session

DATABASE_URI = "postgresql://postgres:password@localhost/testdb2"
engine = create_engine(DATABASE_URI, connect_args={}, future=True, echo=True)


def get_session():
    with Session(engine) as session:
        yield session
"""


from sqlalchemy import create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker
from sqlalchemy.ext.asyncio import create_async_engine
#from sqlalchemy.ext.asyncio.session import AsyncSession
from sqlmodel.ext.asyncio.session import AsyncSession 

DB_URI = "postgresql://postgres:password@localhost/trackerdb"
ASYNC_DB_URI = "postgresql+asyncpg://postgres:password@localhost/trackerdb"

engine = create_engine(DB_URI, connect_args={}, future=True)

SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine, future=True)

Base = declarative_base()


# non async session
def get_session():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()


async_engine = create_async_engine(ASYNC_DB_URI, echo=True, future=True)


# == Async session by SQLAlchemy implementation
async def get_async_session() -> AsyncSession:
    async_session = sessionmaker(bind=async_engine, class_=AsyncSession,
                                 expire_on_commit=False)

    async with async_session() as session:
        yield session
