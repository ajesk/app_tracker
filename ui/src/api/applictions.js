const getUserApplications = async (userId) => {
    try {
        const response = await fetch(`http://localhost:8080/user/${userId}/application`);
        return await response.json();
    } catch (error) {
        console.log('Error:', error);
    }
};

export { getUserApplications };