import {PhotoAddSchema} from "@colday/api";
import {Image, Group, Header, SimpleCell, Flex, Button, FormItem} from "@vkontakte/vkui";

type VkAdminFromPhotosDeleteProps = {
    photo?: PhotoAddSchema
    onDelete?: (id: number) => void
}
const VkAdminFormPhotosDelete = ({photo, onDelete}: VkAdminFromPhotosDeleteProps) => {
    return <Group mode="card">
        <SimpleCell indicator={photo?.id}>
            ID
        </SimpleCell>
        <SimpleCell indicator={photo?.filename}>
            Название файла
        </SimpleCell>
        <Group header={<Header mode="secondary">ПРОСМОТР</Header>} mode="plain">
            <Flex margin="auto" direction="column" gap="m">
                <div style={{height: '20vh'}}>
                    <Image keepAspectRatio src={`/api/v1/photos/${photo?.id}`} widthSize="100%"/>
                </div>
            </Flex>
        </Group>
        <FormItem>
            <Button onClick={onDelete && onDelete.bind(onDelete, photo?.id || 0)} mode="tertiary" appearance="negative"
                    size="l" stretched>
                Удалить
            </Button>
        </FormItem>
    </Group>

}

export default VkAdminFormPhotosDelete;
