import http from 'k6/http';
import { check, sleep } from 'k6';
import { randomString } from 'https://jslib.k6.io/k6-utils/1.1.0/index.js';

function getRandomUserType() {
    const userTypes = ['moderator', 'client'];
    return userTypes[Math.floor(Math.random() * userTypes.length)];
}

function getRandomPrice(min, max) {
    return Math.floor(Math.random() * (max - min + 1)) + min;
}

function getRandomRooms(min, max) {
    return Math.floor(Math.random() * (max - min + 1)) + min;
}

export let options = {
    stages: [
        { duration: '1m', target: 1000 },
        { duration: '2m', target: 1000 },
        { duration: '1m', target: 0 },
    ],
    thresholds: {
        http_req_duration: ['p(99)<50'],
        http_req_failed: ['rate<0.0001'],
    },
};

export default function () {
    // Registration
    const email = `testuser${randomString(5)}@example.com`;
    const password = `testpassword${randomString(10)}`;
    const userType = getRandomUserType();

    let registrationData = JSON.stringify({
        email: email,
        password: password,
        user_type: userType
    });

    let res = http.post('http://localhost:8080/auth/register', registrationData, {
        headers: {'Content-Type': 'application/json'},
    });
    check(res, {'register status 200': (r) => r.status === 200});

    sleep(1);

    // Login
    let userID = res.json().userID;

    let loginData = JSON.stringify({
        userID: userID,
        password: password
    });

    res = http.post('http://localhost:8080/auth/login', loginData, {
        headers: {'Content-Type': 'application/json'},
    });
    check(res, {'login status 200': (r) => r.status === 200});

    let authToken = res.json().token;

    sleep(1);

    // Create House
    if (userType === 'moderator') {
        let houseData = JSON.stringify({
            address: `Test Address ${randomString(10)}`,
            year: 2024,
            developer: `Test Developer ${randomString(10)}`
        });

        res = http.post('http://localhost:8080/realty/house/create', houseData, {
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${authToken}`
            },
        });
        check(res, {'create house status 200': (r) => r.status === 200});

        let houseID = res.json().id;

        sleep(1);
    }

    // Create Flat
    let flatData = JSON.stringify({
        houseID: houseID,
        price: getRandomPrice(2000000, 100000000),
        rooms: getRandomRooms(1, 20)
    });

    res = http.post('http://localhost:8080/realty/flat/create', flatData, {
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${authToken}`
        },
    });
    check(res, {'create flat status 200': (r) => r.status === 200});

    let flatID = res.json().id;

    sleep(1);

    // Update Flat Status
    let flatStatus = res.json().status;

    if (flatStatus === 'created') {
        let updateFlatStatusDataOne = JSON.stringify({
            flatID: flatID,
            newStatus: 'on moderation'
        });

        res = http.patch('http://localhost:8080/realty/flat/update', updateFlatStatusDataOne, {
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${authToken}`
            },
        });
        check(res, {'update flat status to on moderation 200': (r) => r.status === 200});

        sleep(1);

        res = http.get(`http://localhost:8080/realty/flat/${flatID}`, {
            headers: {'Authorization': `Bearer ${authToken}`}
        });
        check(res, {'get flat status 200 after update': (r) => r.status === 200});

        flatStatus = res.json().status;

        if (flatStatus === 'on moderation') {
            let updateFlatStatusDataTwo = JSON.stringify({
                flatID: flatID,
                newStatus: 'approved'
            });

            res = http.patch('http://localhost:8080/realty/flat/update', updateFlatStatusDataTwo, {
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${authToken}`
                },
            });
            check(res, {'update flat status to approved 200': (r) => r.status === 200});

            sleep(1);
        }
    }

    // House Subscribe
    let subscribeData = JSON.stringify({
        email: email
    });

    res = http.post(`http://localhost:8080/house/${houseID}/subscribe`, subscribeData, {
        headers: {'Content-Type': 'application/json'},
    });
    check(res, {'subscribe to house status 200': (r) => r.status === 200});

    sleep(1);
}