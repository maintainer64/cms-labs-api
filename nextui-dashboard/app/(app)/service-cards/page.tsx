import React from 'react';
import { Layout } from '@/components/layout/layout';
import { ServiceCardsEdit } from '@/components/pages/service-cards/edit/service-cards-edit';
import { ServiceCardsList } from '@/components/pages/service-cards';

export const ServiceCardsPage = () => {
  return (
    <Layout>
      <ServiceCardsList />
    </Layout>
  );
};

export const ServiceCardsPageEdit = () => {
  return (
    <Layout>
      <ServiceCardsEdit />
    </Layout>
  );
};
