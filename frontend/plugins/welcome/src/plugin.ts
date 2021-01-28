import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import Login from './components/Login';
import Payment from './components/Payment';
import Equipment from './components/Equipment';
import Bookcourse from './components/ฺฺBookcourse';
import Employee from './components/Employee';

import Promotion from './components/Promotion';
import Promotionsearch from './components/Promotion/Promotionsearch';
import Bookcoursesearch from './components/ฺฺBookcourse/Bookcoursesearch';
import Equipmentsearch from './components/Equipment/Equipmentsearch';
import Equipmentrentalsearch from './components/Equipmentrental/Equipmentrentalsearch';


import Equipmentrental from './components/Equipmentrental';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', Login);
    router.registerRoute('/WelcomePage', WelcomePage);
    router.registerRoute('/Payment', Payment);
    router.registerRoute('/Equipment', Equipment);
    router.registerRoute('/Bookcourse', Bookcourse);
    router.registerRoute('/Employee', Employee);

    router.registerRoute('/Promotion', Promotion);
    router.registerRoute('/Promotionsearch', Promotionsearch);
    router.registerRoute('/Bookcoursesearch', Bookcoursesearch);
    router.registerRoute('/Equipmentsearch', Equipmentsearch);
    router.registerRoute('/Equipmentrentalsearch', Equipmentrentalsearch);

    router.registerRoute('/Equipmentrental', Equipmentrental);
  },
});