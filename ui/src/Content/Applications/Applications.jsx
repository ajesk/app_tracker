import { Text } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import { getUserApplications } from "../../api/applictions";

const Applications = () => {
    const [applications, setApplications] = useState([]);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const applications = await getUserApplications(1);
                setApplications(applications);
            } catch (error) {
                console.log('Error:', error);
            }


        }
        fetchData();
    }, []);

    return applications.map((application) => <Text>{application.title}</Text>
    );
};

export default Applications;