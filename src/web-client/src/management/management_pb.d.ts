import * as jspb from 'google-protobuf'

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb'; // proto import: "google/protobuf/empty.proto"


export class ChangePasswordRequest extends jspb.Message {
  getToken(): string;
  setToken(value: string): ChangePasswordRequest;

  getUsername(): string;
  setUsername(value: string): ChangePasswordRequest;

  getOldPassword(): string;
  setOldPassword(value: string): ChangePasswordRequest;

  getNewPassword(): string;
  setNewPassword(value: string): ChangePasswordRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ChangePasswordRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ChangePasswordRequest): ChangePasswordRequest.AsObject;
  static serializeBinaryToWriter(message: ChangePasswordRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ChangePasswordRequest;
  static deserializeBinaryFromReader(message: ChangePasswordRequest, reader: jspb.BinaryReader): ChangePasswordRequest;
}

export namespace ChangePasswordRequest {
  export type AsObject = {
    token: string,
    username: string,
    oldPassword: string,
    newPassword: string,
  }
}

export class ChangePasswordResponse extends jspb.Message {
  getStatus(): Status;
  setStatus(value: Status): ChangePasswordResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ChangePasswordResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ChangePasswordResponse): ChangePasswordResponse.AsObject;
  static serializeBinaryToWriter(message: ChangePasswordResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ChangePasswordResponse;
  static deserializeBinaryFromReader(message: ChangePasswordResponse, reader: jspb.BinaryReader): ChangePasswordResponse;
}

export namespace ChangePasswordResponse {
  export type AsObject = {
    status: Status,
  }
}

export class AuthenticateRequest extends jspb.Message {
  getToken(): string;
  setToken(value: string): AuthenticateRequest;

  getUsername(): string;
  setUsername(value: string): AuthenticateRequest;

  getPassword(): string;
  setPassword(value: string): AuthenticateRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AuthenticateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AuthenticateRequest): AuthenticateRequest.AsObject;
  static serializeBinaryToWriter(message: AuthenticateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AuthenticateRequest;
  static deserializeBinaryFromReader(message: AuthenticateRequest, reader: jspb.BinaryReader): AuthenticateRequest;
}

export namespace AuthenticateRequest {
  export type AsObject = {
    token: string,
    username: string,
    password: string,
  }
}

export class AuthenticateResponse extends jspb.Message {
  getStatus(): Status;
  setStatus(value: Status): AuthenticateResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AuthenticateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AuthenticateResponse): AuthenticateResponse.AsObject;
  static serializeBinaryToWriter(message: AuthenticateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AuthenticateResponse;
  static deserializeBinaryFromReader(message: AuthenticateResponse, reader: jspb.BinaryReader): AuthenticateResponse;
}

export namespace AuthenticateResponse {
  export type AsObject = {
    status: Status,
  }
}

export class UsernameRequest extends jspb.Message {
  getToken(): string;
  setToken(value: string): UsernameRequest;

  getOldUsername(): string;
  setOldUsername(value: string): UsernameRequest;

  getNewUsername(): string;
  setNewUsername(value: string): UsernameRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UsernameRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UsernameRequest): UsernameRequest.AsObject;
  static serializeBinaryToWriter(message: UsernameRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UsernameRequest;
  static deserializeBinaryFromReader(message: UsernameRequest, reader: jspb.BinaryReader): UsernameRequest;
}

export namespace UsernameRequest {
  export type AsObject = {
    token: string,
    oldUsername: string,
    newUsername: string,
  }
}

export class UsernameResponse extends jspb.Message {
  getStatus(): Status;
  setStatus(value: Status): UsernameResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UsernameResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UsernameResponse): UsernameResponse.AsObject;
  static serializeBinaryToWriter(message: UsernameResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UsernameResponse;
  static deserializeBinaryFromReader(message: UsernameResponse, reader: jspb.BinaryReader): UsernameResponse;
}

export namespace UsernameResponse {
  export type AsObject = {
    status: Status,
  }
}

export class DisconnectRequest extends jspb.Message {
  getToken(): string;
  setToken(value: string): DisconnectRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DisconnectRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DisconnectRequest): DisconnectRequest.AsObject;
  static serializeBinaryToWriter(message: DisconnectRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DisconnectRequest;
  static deserializeBinaryFromReader(message: DisconnectRequest, reader: jspb.BinaryReader): DisconnectRequest;
}

export namespace DisconnectRequest {
  export type AsObject = {
    token: string,
  }
}

export class SessionRequest extends jspb.Message {
  getHostname(): string;
  setHostname(value: string): SessionRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SessionRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SessionRequest): SessionRequest.AsObject;
  static serializeBinaryToWriter(message: SessionRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SessionRequest;
  static deserializeBinaryFromReader(message: SessionRequest, reader: jspb.BinaryReader): SessionRequest;
}

export namespace SessionRequest {
  export type AsObject = {
    hostname: string,
  }
}

export class SessionResponse extends jspb.Message {
  getToken(): string;
  setToken(value: string): SessionResponse;

  getStatus(): Status;
  setStatus(value: Status): SessionResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SessionResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SessionResponse): SessionResponse.AsObject;
  static serializeBinaryToWriter(message: SessionResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SessionResponse;
  static deserializeBinaryFromReader(message: SessionResponse, reader: jspb.BinaryReader): SessionResponse;
}

export namespace SessionResponse {
  export type AsObject = {
    token: string,
    status: Status,
  }
}

export class User extends jspb.Message {
  getUsername(): string;
  setUsername(value: string): User;
  hasUsername(): boolean;
  clearUsername(): User;

  getPassword(): string;
  setPassword(value: string): User;
  hasPassword(): boolean;
  clearPassword(): User;

  getCommonName(): string;
  setCommonName(value: string): User;
  hasCommonName(): boolean;
  clearCommonName(): User;

  getFirstName(): string;
  setFirstName(value: string): User;
  hasFirstName(): boolean;
  clearFirstName(): User;

  getLastName(): string;
  setLastName(value: string): User;
  hasLastName(): boolean;
  clearLastName(): User;

  getEmployeeNumber(): string;
  setEmployeeNumber(value: string): User;
  hasEmployeeNumber(): boolean;
  clearEmployeeNumber(): User;

  getPhoneNumber(): string;
  setPhoneNumber(value: string): User;
  hasPhoneNumber(): boolean;
  clearPhoneNumber(): User;

  getAddress(): string;
  setAddress(value: string): User;
  hasAddress(): boolean;
  clearAddress(): User;

  getCompanyRole(): string;
  setCompanyRole(value: string): User;
  hasCompanyRole(): boolean;
  clearCompanyRole(): User;

  getTeam(): string;
  setTeam(value: string): User;
  hasTeam(): boolean;
  clearTeam(): User;

  getReportsTo(): string;
  setReportsTo(value: string): User;
  hasReportsTo(): boolean;
  clearReportsTo(): User;

  getOtherCharacteristicsMap(): jspb.Map<string, string>;
  clearOtherCharacteristicsMap(): User;

  getMemberOfList(): Array<string>;
  setMemberOfList(value: Array<string>): User;
  clearMemberOfList(): User;
  addMemberOf(value: string, index?: number): User;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): User.AsObject;
  static toObject(includeInstance: boolean, msg: User): User.AsObject;
  static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): User;
  static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
  export type AsObject = {
    username?: string,
    password?: string,
    commonName?: string,
    firstName?: string,
    lastName?: string,
    employeeNumber?: string,
    phoneNumber?: string,
    address?: string,
    companyRole?: string,
    team?: string,
    reportsTo?: string,
    otherCharacteristicsMap: Array<[string, string]>,
    memberOfList: Array<string>,
  }

  export enum UsernameCase { 
    _USERNAME_NOT_SET = 0,
    USERNAME = 1,
  }

  export enum PasswordCase { 
    _PASSWORD_NOT_SET = 0,
    PASSWORD = 2,
  }

  export enum CommonNameCase { 
    _COMMON_NAME_NOT_SET = 0,
    COMMON_NAME = 3,
  }

  export enum FirstNameCase { 
    _FIRST_NAME_NOT_SET = 0,
    FIRST_NAME = 4,
  }

  export enum LastNameCase { 
    _LAST_NAME_NOT_SET = 0,
    LAST_NAME = 5,
  }

  export enum EmployeeNumberCase { 
    _EMPLOYEE_NUMBER_NOT_SET = 0,
    EMPLOYEE_NUMBER = 6,
  }

  export enum PhoneNumberCase { 
    _PHONE_NUMBER_NOT_SET = 0,
    PHONE_NUMBER = 7,
  }

  export enum AddressCase { 
    _ADDRESS_NOT_SET = 0,
    ADDRESS = 8,
  }

  export enum CompanyRoleCase { 
    _COMPANY_ROLE_NOT_SET = 0,
    COMPANY_ROLE = 9,
  }

  export enum TeamCase { 
    _TEAM_NOT_SET = 0,
    TEAM = 10,
  }

  export enum ReportsToCase { 
    _REPORTS_TO_NOT_SET = 0,
    REPORTS_TO = 11,
  }
}

export class UserRequest extends jspb.Message {
  getToken(): string;
  setToken(value: string): UserRequest;

  getUsername(): string;
  setUsername(value: string): UserRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UserRequest): UserRequest.AsObject;
  static serializeBinaryToWriter(message: UserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserRequest;
  static deserializeBinaryFromReader(message: UserRequest, reader: jspb.BinaryReader): UserRequest;
}

export namespace UserRequest {
  export type AsObject = {
    token: string,
    username: string,
  }
}

export class UserResponse extends jspb.Message {
  getUser(): User | undefined;
  setUser(value?: User): UserResponse;
  hasUser(): boolean;
  clearUser(): UserResponse;

  getStatus(): Status;
  setStatus(value: Status): UserResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UserResponse): UserResponse.AsObject;
  static serializeBinaryToWriter(message: UserResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserResponse;
  static deserializeBinaryFromReader(message: UserResponse, reader: jspb.BinaryReader): UserResponse;
}

export namespace UserResponse {
  export type AsObject = {
    user?: User.AsObject,
    status: Status,
  }

  export enum UserCase { 
    _USER_NOT_SET = 0,
    USER = 1,
  }
}

export class DeleteUsersResponse extends jspb.Message {
  getStatus(): Status;
  setStatus(value: Status): DeleteUsersResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteUsersResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteUsersResponse): DeleteUsersResponse.AsObject;
  static serializeBinaryToWriter(message: DeleteUsersResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteUsersResponse;
  static deserializeBinaryFromReader(message: DeleteUsersResponse, reader: jspb.BinaryReader): DeleteUsersResponse;
}

export namespace DeleteUsersResponse {
  export type AsObject = {
    status: Status,
  }
}

export class AddUsersRequest extends jspb.Message {
  getToken(): string;
  setToken(value: string): AddUsersRequest;

  getUsersList(): Array<User>;
  setUsersList(value: Array<User>): AddUsersRequest;
  clearUsersList(): AddUsersRequest;
  addUsers(value?: User, index?: number): User;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddUsersRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AddUsersRequest): AddUsersRequest.AsObject;
  static serializeBinaryToWriter(message: AddUsersRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddUsersRequest;
  static deserializeBinaryFromReader(message: AddUsersRequest, reader: jspb.BinaryReader): AddUsersRequest;
}

export namespace AddUsersRequest {
  export type AsObject = {
    token: string,
    usersList: Array<User.AsObject>,
  }
}

export class AddUsersResponse extends jspb.Message {
  getStatus(): Status;
  setStatus(value: Status): AddUsersResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddUsersResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AddUsersResponse): AddUsersResponse.AsObject;
  static serializeBinaryToWriter(message: AddUsersResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddUsersResponse;
  static deserializeBinaryFromReader(message: AddUsersResponse, reader: jspb.BinaryReader): AddUsersResponse;
}

export namespace AddUsersResponse {
  export type AsObject = {
    status: Status,
  }
}

export class DeleteUsersRequest extends jspb.Message {
  getToken(): string;
  setToken(value: string): DeleteUsersRequest;

  getUsernamesList(): Array<string>;
  setUsernamesList(value: Array<string>): DeleteUsersRequest;
  clearUsernamesList(): DeleteUsersRequest;
  addUsernames(value: string, index?: number): DeleteUsersRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteUsersRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteUsersRequest): DeleteUsersRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteUsersRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteUsersRequest;
  static deserializeBinaryFromReader(message: DeleteUsersRequest, reader: jspb.BinaryReader): DeleteUsersRequest;
}

export namespace DeleteUsersRequest {
  export type AsObject = {
    token: string,
    usernamesList: Array<string>,
  }
}

export class ModifyUsersRequest extends jspb.Message {
  getToken(): string;
  setToken(value: string): ModifyUsersRequest;

  getUsernamesList(): Array<string>;
  setUsernamesList(value: Array<string>): ModifyUsersRequest;
  clearUsernamesList(): ModifyUsersRequest;
  addUsernames(value: string, index?: number): ModifyUsersRequest;

  getFilterList(): Array<Filter>;
  setFilterList(value: Array<Filter>): ModifyUsersRequest;
  clearFilterList(): ModifyUsersRequest;
  addFilter(value?: Filter, index?: number): Filter;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ModifyUsersRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ModifyUsersRequest): ModifyUsersRequest.AsObject;
  static serializeBinaryToWriter(message: ModifyUsersRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ModifyUsersRequest;
  static deserializeBinaryFromReader(message: ModifyUsersRequest, reader: jspb.BinaryReader): ModifyUsersRequest;
}

export namespace ModifyUsersRequest {
  export type AsObject = {
    token: string,
    usernamesList: Array<string>,
    filterList: Array<Filter.AsObject>,
  }
}

export class ModifyUsersResponse extends jspb.Message {
  getStatus(): Status;
  setStatus(value: Status): ModifyUsersResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ModifyUsersResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ModifyUsersResponse): ModifyUsersResponse.AsObject;
  static serializeBinaryToWriter(message: ModifyUsersResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ModifyUsersResponse;
  static deserializeBinaryFromReader(message: ModifyUsersResponse, reader: jspb.BinaryReader): ModifyUsersResponse;
}

export namespace ModifyUsersResponse {
  export type AsObject = {
    status: Status,
  }
}

export class ListUsersRequest extends jspb.Message {
  getToken(): string;
  setToken(value: string): ListUsersRequest;

  getUsername(): string;
  setUsername(value: string): ListUsersRequest;
  hasUsername(): boolean;
  clearUsername(): ListUsersRequest;

  getFilterList(): Array<Filter>;
  setFilterList(value: Array<Filter>): ListUsersRequest;
  clearFilterList(): ListUsersRequest;
  addFilter(value?: Filter, index?: number): Filter;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListUsersRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListUsersRequest): ListUsersRequest.AsObject;
  static serializeBinaryToWriter(message: ListUsersRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListUsersRequest;
  static deserializeBinaryFromReader(message: ListUsersRequest, reader: jspb.BinaryReader): ListUsersRequest;
}

export namespace ListUsersRequest {
  export type AsObject = {
    token: string,
    username?: string,
    filterList: Array<Filter.AsObject>,
  }

  export enum UsernameCase { 
    _USERNAME_NOT_SET = 0,
    USERNAME = 2,
  }
}

export class Filter extends jspb.Message {
  getCharacteristic(): Characteristic;
  setCharacteristic(value: Characteristic): Filter;

  getValue(): string;
  setValue(value: string): Filter;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Filter.AsObject;
  static toObject(includeInstance: boolean, msg: Filter): Filter.AsObject;
  static serializeBinaryToWriter(message: Filter, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Filter;
  static deserializeBinaryFromReader(message: Filter, reader: jspb.BinaryReader): Filter;
}

export namespace Filter {
  export type AsObject = {
    characteristic: Characteristic,
    value: string,
  }
}

export class ListUsersResponse extends jspb.Message {
  getUsersList(): Array<User>;
  setUsersList(value: Array<User>): ListUsersResponse;
  clearUsersList(): ListUsersResponse;
  addUsers(value?: User, index?: number): User;

  getStatus(): Status;
  setStatus(value: Status): ListUsersResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListUsersResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListUsersResponse): ListUsersResponse.AsObject;
  static serializeBinaryToWriter(message: ListUsersResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListUsersResponse;
  static deserializeBinaryFromReader(message: ListUsersResponse, reader: jspb.BinaryReader): ListUsersResponse;
}

export namespace ListUsersResponse {
  export type AsObject = {
    usersList: Array<User.AsObject>,
    status: Status,
  }
}

export enum Status { 
  STATUS_OK = 0,
  STATUS_ERROR = 1,
  STATUS_USER_NOT_FOUND = 2,
}
export enum Characteristic { 
  COMMON_NAME = 0,
  FIRST_NAME = 1,
  LAST_NAME = 2,
  EMPLOYEE_NUMBER = 3,
  PHONE_NUMBER = 4,
  ADDRESS = 5,
  COMPANY_ROLE = 6,
  TEAM = 7,
  REPORTS_TO = 8,
  OTHER = 9,
  MEMBER_OF = 10,
  USERNAME = 11,
}
