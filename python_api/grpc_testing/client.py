from concurrent import futures

import grpc
from protos import users_pb2
from protos import users_pb2_grpc

from google.protobuf.json_format import MessageToDict

URI = "localhost:50051"


def run(URI):
    with grpc.insecure_channel(URI) as channel:
        stub = users_pb2_grpc.UsersStub(channel)

        get_users_responses = stub.GetUsers(users_pb2.GetUsersRequest())
        print(get_users_responses)
        for response in get_users_responses:
            print(response)


        """
        get_user_by_id_response = stub.GetUserById(users_pb2.GetUserByIdRequest(id="1"))

        u3 = users_pb2.User(id="3", name="Funky", email="mail.as", password="pwwpww")
        create_user_response = stub.CreateUser(users_pb2.CreateUserRequest(user=u3))

        u3 = users_pb2.User(id="3", name="Crap", email="mail.as", password="pwwpww")
        update_user_response = stub.UpdateUser(users_pb2.UpdateUserRequest(user=u3))

        get_another_users_response = stub.GetUsers(users_pb2.GetUsersRequest())

        delete_user_response = stub.DeleteUser(users_pb2.DeleteUserRequest(user=u3))

        get_another2_users_response = stub.GetUsers(users_pb2.GetUsersRequest())
        """

    """

    user = get_users_response.users[0]
    test = MessageToDict(user)

    print(test)
    print(type(test))
    print("-------------------")

    print(get_user_by_id_response)
    print(type(get_user_by_id_response))

    print("create new users--")
    print(create_user_response)
    print(type(create_user_response))

    print("--update users ---")
    print(update_user_response)
    print(type(update_user_response))

    print("--check users again--")
    print(get_another_users_response.users)
    print(type(get_another_users_response.users))
    print(type(get_another_users_response))

    print("--deleted--")
    print(delete_user_response)

    print("---after--delete--")
    print(get_another2_users_response)

    """

if __name__ == "__main__":
    run(URI)
