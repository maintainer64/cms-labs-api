import { CamelCasedPropertiesDeep } from 'type-fest';
import { UsecasesTargetItem, UsecasesTargetRelationItem } from '@/helpers/api';

export type TargetNodeType = CamelCasedPropertiesDeep<UsecasesTargetItem>;
export type TargetRelationType = CamelCasedPropertiesDeep<UsecasesTargetRelationItem>;

export const filterTargets = (
  targets: TargetNodeType[],
  searchQuery: string = '',
  filter: string = 'all'
): TargetNodeType[] => {
  if (!targets) return [];

  // Фильтрация по поисковому запросу
  const searchFiltered = searchQuery.trim()
    ? targets.filter((target) =>
        [
          target.taget.name?.toLowerCase(),
          target.taget.id?.toLowerCase(),
          target.taget.description?.toLowerCase() || ''
        ].some((field) => field?.includes(searchQuery.toLowerCase()))
      )
    : targets;

  if (filter === 'my') {
    return searchFiltered.filter((target) => target?.isMine);
  }
  if (filter === 'all') {
    return searchFiltered;
  }

  return searchFiltered.filter(
    (target) =>
      (target?.taget?.tags as any)?.['tags'].include(filter) ||
      (target?.taget?.internalTags as any)?.['tags']?.include(filter)
  );
};
