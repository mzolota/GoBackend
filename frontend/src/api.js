import axios from "axios";

const apiClient = axios.create({
    baseURL: 'http://localhost:8080',  // backend URL
    headers: {
        'Content-Type': 'application/json',
    },
});

export function loginUser(credentials) {
    // POST za prijavu
    return apiClient.post('/login', credentials);
}

export function registerUser(userData) {
    // POST za registraciju
    return apiClient.post('/users', userData);
}
