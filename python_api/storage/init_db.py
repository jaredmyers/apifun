from sqlalchemy import create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker

DATABASE_URI = "postgresql://postgres@localhost/testdb"

engine = create_engine(DATABASE_URI, connect_args={}, future=True)

SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine, future=True)

Base = declarative_base()


def get_session():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()
