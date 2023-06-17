import config from "../config";

const getUserApplications = async (userId) => {
    try {
        const response = await fetch(`${config.url}/user/${userId}/application`);
        return await response.json();
    } catch (error) {
        console.error('error caught while getting user applications: ', error);
    }
};

const getApplicationById = async (applicationId) => {
    try {
        const response = await fetch(`${config.url}/application/${applicationId}`);
        return await response.json();
    } catch (error) {
        console.error('error caught while getting applications: ', error);
    }
}

const createApplication = async (userId, companyId, application) => {
    try {
        fetch(`${config.url}/user/${userId}/company/${companyId}/application`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(application)
        })
            .then((response) => response.json())
            .then((data) => {
                // Handle the response data
                console.log(data);
            })
            .catch((error) => {
                // Handle the error
                console.error(error);
            });

    } catch (error) {
        console.error('error caught while posting application: ', error);
    }
}

export { getUserApplications, getApplicationById, createApplication };