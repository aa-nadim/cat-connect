// js/components/Homepage.js
class Homepage {
    constructor() {
        this.header = new Header();
        this.tabContent = new TabContent();
    }

    render(activeTab) {
        const app = document.getElementById('app');
        app.innerHTML = `
            <div class="min-h-screen bg-white">
                <div class="container">
                    ${this.header.render()}
                    <main class="row py-4">
                        <div class="col-lg-6 mb-4 mb-lg-0">
                            <h1 class="display-4 fw-bold mb-3">
                                The Cat API
                                <br>
                                <span class="text-orange-500">Cats as a service.</span>
                            </h1>
                            <p class="lead mb-3">Because everyday is a Caturday.</p>
                            <p class="mb-4">An API all about cats.<br>60k+ Images. Breeds. Facts.</p>
                            <div class="d-flex flex-column flex-sm-row gap-3">
                                <button class="btn btn-dark">GET YOUR API KEY</button>
                                <button class="btn btn-outline-dark">READ OUR GUIDES</button>
                            </div>
                        </div>
                        <div class="col-lg-6">
                            ${this.tabContent.render(activeTab)}
                        </div>
                    </main>
                </div>
            </div>
        `;
    }
}
