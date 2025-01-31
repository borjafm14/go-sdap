import { grpc } from '@improbable-eng/grpc-web';
import { managementService } from '../proto/management_pb_service';
import { SessionRequest, AuthenticateRequest, ListUsersRequest, AddUsersRequest, DeleteUsersRequest } from '../proto/management_pb';

const MANAGEMENT_SERVER_URL = 'http://localhost:50052'; // Update with your management server URL

export const login = (username, password) => {
    return new Promise((resolve, reject) => {
        const request = new AuthenticateRequest();
        request.setUsername(username);
        request.setPassword(password);

        grpc.invoke(managementService.Authenticate, {
            request: request,
            host: MANAGEMENT_SERVER_URL,
            onMessage: (response) => {
                resolve(response.getStatus());
            },
            onEnd: (code, msg) => {
                if (code !== grpc.Code.OK) {
                    reject(new Error(msg));
                }
            }
        });
    });
};

export const fetchUsers = (token) => {
    return new Promise((resolve, reject) => {
        const request = new ListUsersRequest();
        request.setToken(token);

        grpc.invoke(managementService.ListUsers, {
            request: request,
            host: MANAGEMENT_SERVER_URL,
            onMessage: (response) => {
                resolve(response.getUsersList());
            },
            onEnd: (code, msg) => {
                if (code !== grpc.Code.OK) {
                    reject(new Error(msg));
                }
            }
        });
    });
};

export const createUser = (token, user) => {
    return new Promise((resolve, reject) => {
        const request = new AddUsersRequest();
        request.setToken(token);
        request.setUsersList([user]);

        grpc.invoke(managementService.AddUsers, {
            request: request,
            host: MANAGEMENT_SERVER_URL,
            onMessage: (response) => {
                resolve(response.getStatus());
            },
            onEnd: (code, msg) => {
                if (code !== grpc.Code.OK) {
                    reject(new Error(msg));
                }
            }
        });
    });
};

export const deleteUser = (token, usernames) => {
    return new Promise((resolve, reject) => {
        const request = new DeleteUsersRequest();
        request.setToken(token);
        request.setUsernamesList(usernames);

        grpc.invoke(managementService.DeleteUsers, {
            request: request,
            host: MANAGEMENT_SERVER_URL,
            onMessage: (response) => {
                resolve(response.getStatus());
            },
            onEnd: (code, msg) => {
                if (code !== grpc.Code.OK) {
                    reject(new Error(msg));
                }
            }
        });
    });
};