import { Text } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import { getUserApplications } from "../../api/applictions";
import ApplicationTable from "./ApplicationTable";

const Applications = () => {
    const [applications, setApplications] = useState([]);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const res = await getUserApplications(1);
                setApplications(res);
            } catch (error) {
                console.error('Error:', error);
            }
        }
        fetchData();
    }, []);

    console.log(applications)
    return <ApplicationTable applications={applications} />;
};

export default Applications;