// static/js/router.js
class Router {
    constructor() {
        this.routes = {
            '/': 'voting',
            '/voting': 'voting',
            '/breeds': 'breeds',
            '/favs': 'favs'
        };

        window.addEventListener('popstate', () => {
            this.handleRoute();
        });
    }

    navigate(path) {
        if (path === '/') path = '/voting';
        window.history.pushState({}, '', path);
        this.handleRoute();
    }

    handleRoute() {
        const path = window.location.pathname;
        const activeTab = this.routes[path] || 'voting';
        this.renderHomepage(activeTab);
    }

    renderHomepage(activeTab) {
        const homepage = new Homepage();
        homepage.render(activeTab);
        
        // Initialize the active tab without reloading other tabs
        switch(activeTab) {
            case 'voting':
                if (!votingTab.initialized) {
                    votingTab.init();
                }
                break;
            case 'breeds':
                if (!breedsTab.initialized) {
                    breedsTab.init();
                }
                break;
            case 'favs':
                if (!favsTab.initialized) {
                    favsTab.init();
                }
                break;
        }
    }
}