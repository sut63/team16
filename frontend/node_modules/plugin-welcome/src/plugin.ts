import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import Login from './components/Login';
import Payment from './components/Payment';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', Login);
    router.registerRoute('/WelcomePage', WelcomePage);
    router.registerRoute('/Payment', Payment);
  },
});