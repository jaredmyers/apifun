from fastapi import APIRouter, Depends
from storage.initdb import get_session, get_async_session
from storage.models import User, Store, Food
from sqlmodel import select 
from fastapi.responses import JSONResponse
from sqlalchemy.orm import selectinload, joinedload

router = APIRouter()


# == Store Routes ==

@router.get("/stores")
async def get_stores(session=Depends(get_async_session)):
    statement = select(Store)
    stores = await session.execute(statement)
    return stores.all()


@router.get("/stores/{id}")
async def get_store_by_id(session=Depends(get_async_session)):
    statement = select(Store).where(Store.id == id)
    result = await session.execute(statement)
    return result.first()


# create, update, delete store
@router.post("/stores")
async def create_store(session=Depends(get_async_session)):
    pass


@router.delete("/stores/{id}")
async def delete_store(session=Depends(get_async_session)):
    pass


@router.put("/stores/{id}")
async def update_store(session=Depends(get_async_session)):
    pass
