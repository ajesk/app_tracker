
import React from 'react';
import { Table, Thead, Tbody, Tr, Th, TableContainer } from '@chakra-ui/react';
import ApplicationRow from './ApplicationRow';

const ApplicationTable = ({ applications = [] }) => {
    return (
        <TableContainer>
            <Table variant="simple" colorScheme="gray">
                <Thead>
                    <Tr>
                        <Th>Title</Th>
                        <Th>Company</Th>
                    </Tr>
                </Thead>
                <Tbody>
                    {
                        applications.map((application) => <ApplicationRow application={application} />)
                    }
                </Tbody>
            </Table>
        </TableContainer>
    );
};

export default ApplicationTable;
