from fastapi import APIRouter


router = APIRouter()


# == Users Endpoint
@router.get("/users")
async def get_users():
    pass


@router.post("/users")
async def create_user(user: str):
    pass


@router.get("/user/{id}")
def get_user(user_id: str):
    pass
