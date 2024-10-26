import {Button, FormItem, Group, Input, NativeSelect} from "@vkontakte/vkui";
import {useForm} from "@tanstack/react-form";
import AdminGroupSpiner from "./spiner";
import GroupAlerts from "../alerts/groupAlerts";
import {type GetPhotoListApiV1PhotosListGetResponse, ReleasesSchemaItem} from "@colday/api";

type VkAdminFormReleaseEditProps = {
    values?: ReleasesSchemaItem
    photos?: GetPhotoListApiV1PhotosListGetResponse
    isLoading?: boolean
    onSubmit: (values: ReleasesSchemaItem) => void
    onDelete: (id: number) => void
}

const emptyValues: ReleasesSchemaItem = {
    id: undefined,
    created_at: '',
    updated_at: '',
    description: '',
    file_id: '',
    link: ''
}

const VkAdminFormReleaseEdit = ({values, photos, isLoading, onSubmit, onDelete}: VkAdminFormReleaseEditProps) => {
    const form = useForm({
        defaultValues: values || emptyValues,
        onSubmit: ({value}) => {
            onSubmit(value)
        }
    })
    if (isLoading) return <AdminGroupSpiner/>
    const options = photos?.photos?.map((photo) => {
        return <option key={`option-photo-url-${photo.id}`}
                       value={`/api/v1/photos/${photo.id}`}>
            {photo.id} {photo.filename}
        </option>
    });
    return <Group mode="card">
        <GroupAlerts/>
        <form
            onSubmit={(e) => {
                e.preventDefault()
                e.stopPropagation()
                form.handleSubmit()
            }}
        >
            <form.Field
                name="id"
                children={(field) => {
                    return (
                        <FormItem top="ID" htmlFor={field.name}>
                            <Input
                                id={field.name}
                                disabled={true}
                                type="number"
                                placeholder=""
                                value={field.state.value || ''}
                                onBlur={field.handleBlur}
                            />
                        </FormItem>
                    )
                }}
            />
            <form.Field
                name="description"
                children={(field) => {
                    return (
                        <FormItem top="Описание релиза" htmlFor={field.name}>
                            <Input
                                id={field.name}
                                type="text"
                                placeholder=""
                                value={field.state.value || ''}
                                onBlur={field.handleBlur}
                                onChange={(e) => field.handleChange(e.target.value)}
                            />
                        </FormItem>
                    )
                }}
            />
            <form.Field
                name="link"
                children={(field) => {
                    return (
                        <FormItem top="Ссылка на релиз" htmlFor={field.name}>
                            <Input
                                id={field.name}
                                type="url"
                                placeholder=""
                                value={field.state.value || ''}
                                onBlur={field.handleBlur}
                                onChange={(e) => field.handleChange(e.target.value)}
                            />
                        </FormItem>
                    )
                }}
            />
            <form.Field
                name="file_id"
                children={(field) => {
                    return (
                        <FormItem top="Ссылка на картинку" htmlFor={field.name}>
                            <NativeSelect
                                id={field.name}
                                placeholder="Введите адрес картинки"
                                onChange={
                                    (e) =>
                                        field.handleChange(e.target.value)
                                }
                                value={field.state.value || ''}
                                onBlur={field.handleBlur}
                            >
                                {options}
                            </NativeSelect>
                        </FormItem>
                    )
                }}
            />
            <FormItem>
                <Button type="submit" size="l" stretched>
                    Сохранить
                </Button>
            </FormItem>
        </form>
    </Group>

}

export default VkAdminFormReleaseEdit;
