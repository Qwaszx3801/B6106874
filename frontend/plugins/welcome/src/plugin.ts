import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import Create from './components/Create';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/Create', Create);
  },
});
