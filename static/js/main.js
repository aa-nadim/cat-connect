// static/js/main.js
let votingTab, breedsTab, favsTab, tabContent, header, router;

document.addEventListener('DOMContentLoaded', () => {
    // Initialize components
    header = new Header();
    votingTab = new VotingTab();
    breedsTab = new BreedsTab();
    favsTab = new FavsTab();
    tabContent = new TabContent();
    
    // Initialize router
    router = new Router();
    window.router = router;

    // Handle initial route based on current URL
    const path = window.location.pathname;
    if (path === '/' || path === '') {
        router.navigate('/voting');
    } else {
        router.handleRoute();
    }
});