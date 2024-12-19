document.addEventListener('DOMContentLoaded', function() {
    // Get all view buttons and content containers
    const viewButtons = document.querySelectorAll('[data-view]');
    const viewContents = document.querySelectorAll('.view-content');

    // Function to switch views
    function switchView(viewName) {
        // Update button states
        viewButtons.forEach(btn => {
            btn.classList.toggle('active', btn.dataset.view === viewName);
        });

        // Update view visibility
        viewContents.forEach(content => {
            content.classList.toggle('active', content.id === `${viewName}View`);
        });

        // Trigger content refresh for the active view
        switch(viewName) {
            case 'voting':
                // Refresh voting content
                if (typeof fetchImages === 'function') {
                    fetchImages();
                }
                break;
            case 'favs':
                // Refresh favorites
                if (typeof displayFavorites === 'function') {
                    displayFavorites('grid');
                }
                break;
            case 'breeds':
                // Breeds view is loaded by default
                break;
        }
    }

    // Add click handlers to view buttons
    viewButtons.forEach(button => {
        button.addEventListener('click', () => {
            switchView(button.dataset.view);
        });
    });
});