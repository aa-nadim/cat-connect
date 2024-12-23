// static/js/components/TabContent.js
class TabContent {
    constructor() {
        this.votingTab = votingTab; // Use global instances
        this.breedsTab = breedsTab;
        this.favsTab = favsTab;
    }

    render(activeTab = 'voting') {
        // When the 'favs' tab is selected, fetch the latest favorites
        if (activeTab === 'favs') {
            this.favsTab.fetchFavorites();
        }

        return `
            <div class="card">
                <div class="card-header bg-white">
                    <ul class="nav nav-tabs card-header-tabs">
                        <li class="nav-item">
                            <a class="nav-link ${activeTab === 'voting' ? 'active' : ''}" 
                                onclick="router.navigate('/voting'); return false;">
                                <i class="fas fa-arrow-up-down me-2"></i>Voting
                            </a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link ${activeTab === 'breeds' ? 'active' : ''}" 
                                onclick="router.navigate('/breeds'); return false;">
                                <i class="fas fa-search me-2"></i>Breeds
                            </a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link ${activeTab === 'favs' ? 'active' : ''}" 
                                onclick="router.navigate('/favs'); return false;">
                                <i class="fas fa-heart me-2"></i>Favs
                            </a>
                        </li>
                    </ul>
                </div>
                <div class="card-body" id="tab-content">
                    ${this.getTabContent(activeTab)}
                </div>
            </div>
        `;
    }

    getTabContent(activeTab) {
        switch (activeTab) {
            case 'voting':
                return votingTab.render();
            case 'breeds':
                return breedsTab.render();
            case 'favs':
                return favsTab.render();
            default:
                return votingTab.render();
        }
    }
}