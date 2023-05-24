import { FormControl, FormLabel, Input, Textarea, Button } from '@chakra-ui/react';
import { useState } from 'react';

const CreateApplicationForm = ({ handleSubmit }) => {
    const [formData, setFormData] = useState({
        title: '',
        company: 1,
        notes: '',
        min_salary: 0,
        max_salary: 0
    });

    const handleChange = (e) => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
    };

    const submit = (e) => {
        e.preventDefault();
        handleSubmit(e, formData);
    };

    return (
        <form onSubmit={submit}>
            <FormControl id="title" isRequired>
                <FormLabel>Title</FormLabel>
                <Input type="text" name="title" value={formData.title} onChange={handleChange} />
            </FormControl>
            <FormControl id="company" isRequired>
                <FormLabel>Company</FormLabel>
                <Input type="number" name="company" value={formData.company} onChange={handleChange} />
            </FormControl>
            <FormControl id="min_salary">
                <FormLabel>Min Salary</FormLabel>
                <Input type="number" name="min_salary" value={formData.min_salary} onChange={handleChange} />
            </FormControl>
            <FormControl id="max_salary">
                <FormLabel>Max Salary</FormLabel>
                <Input type="number" name="max_salary" value={formData.max_salary} onChange={handleChange} />
            </FormControl>
            <FormControl id="notes">
                <FormLabel>Notes</FormLabel>
                <Textarea name="notes" value={formData.notes} onChange={handleChange} />
            </FormControl>

            <Button type="submit" colorScheme="blue" mt={4}>
                Submit
            </Button>
        </form>
    );
}

export default CreateApplicationForm;