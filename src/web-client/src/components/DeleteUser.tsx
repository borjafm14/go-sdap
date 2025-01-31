import React, { useState, useEffect } from 'react';
import { deleteUser, getUserList } from '../services/managementService';

const DeleteUser: React.FC = () => {
    const [users, setUsers] = useState<any[]>([]);
    const [selectedUser, setSelectedUser] = useState<string | null>(null);
    const [message, setMessage] = useState<string | null>(null);

    useEffect(() => {
        const fetchUsers = async () => {
            const userList = await getUserList();
            setUsers(userList);
        };

        fetchUsers();
    }, []);

    const handleDelete = async () => {
        if (selectedUser) {
            const response = await deleteUser(selectedUser);
            if (response.status === 'STATUS_OK') {
                setMessage('User deleted successfully');
                setUsers(users.filter(user => user.username !== selectedUser));
            } else {
                setMessage('Failed to delete user');
            }
        }
    };

    return (
        <div>
            <h2>Delete User</h2>
            {message && <p>{message}</p>}
            <select onChange={(e) => setSelectedUser(e.target.value)} value={selectedUser || ''}>
                <option value="" disabled>Select a user to delete</option>
                {users.map(user => (
                    <option key={user.username} value={user.username}>{user.username}</option>
                ))}
            </select>
            <button onClick={handleDelete} disabled={!selectedUser}>Delete User</button>
        </div>
    );
};

export default DeleteUser;