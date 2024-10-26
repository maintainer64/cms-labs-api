import {type ReleasesHandlerApiV1ReleasesGetResponse, ReleasesSchemaItem} from "@colday/api";
import {Avatar, Button, FormItem, Group, Header, SimpleCell} from "@vkontakte/vkui";
import AdminGroupSpiner from "./spiner";
import {AdminPanelMenu} from "@/views/admin/base";
import {Icon28GhostSimpleOutline} from "@vkontakte/icons";
import {useDataProvider} from "../viewParams/hooks";

type VkAdminReleasesListProps = {
    values?: ReleasesHandlerApiV1ReleasesGetResponse
    isLoading?: boolean
    to: (panel: AdminPanelMenu) => void
}

const VkAdminReleasesList = ({values, isLoading, to}: VkAdminReleasesListProps) => {
    const setterView = useDataProvider<ReleasesSchemaItem>();
    if (isLoading) return <AdminGroupSpiner/>
    const components = values?.releases?.map((item) => {
        return <SimpleCell
            before={
                <Avatar
                    size={28}
                    fallbackIcon={<Icon28GhostSimpleOutline/>}
                    src={`/api/v1/photos/${item.file_id}`}
                />
            }
            onClick={() => {
                setterView?.setData?.(item)
                to(AdminPanelMenu.editReleases)
            }} key={item.id} Component="label">
            {item.id}. {item.description}
        </SimpleCell>
    })
    return <Group mode="card">
        <Group header={<Header mode="secondary">ДОБАВИТЬ</Header>} mode="plain">
            <FormItem>
                <Button onClick={() => {
                    setterView?.setData?.({
                        id: null,
                        created_at: new Date().toISOString(),
                        updated_at: new Date().toISOString(),
                        description: '',
                        file_id: '',
                        link: '',
                    })
                    to(AdminPanelMenu.editReleases)
                }} size="l" stretched>
                    Добавить
                </Button>
            </FormItem>
        </Group>
        <Group header={<Header mode="secondary">РЕЛИЗЫ</Header>} mode="plain">
            {components}
        </Group>
    </Group>

}

export default VkAdminReleasesList;
