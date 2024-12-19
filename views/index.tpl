<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Cat Connect</title>
    <link href="/static/css/style.css" rel="stylesheet">
    <link href="/static/css/bootstrap.min.css" rel="stylesheet">

    <link href="/static/css/all.min.css" rel="stylesheet">
</head>
<body>
    <div class="container mt-4">
        <div class="menu-bar">
            <a href="#" class="menu-item bg-warning">
                <i class="fas fa-arrow-up"></i> Voting
            </a>
            <a href="#" class="menu-item active bg-warning">
                <i class="fas fa-search"></i> Breeds
            </a>
            <a href="#" class="menu-item bg-warning">
                <i class="fas fa-heart"></i> Favs
            </a>
        </div>

        <div class="row">
            <div class="col-md-12">
                <h1 class="text-center mb-4">Cat Connect</h1>
                <div class="d-flex justify-content-between mb-4">
                    <div class="breed-selector">
                        <select id="breedSelect" class="form-select">
                            <option value="">Select a breed</option>
                        </select>
                    </div>
                    <div class="navigation-buttons">
                        <button id="prevBtn" class="btn btn-primary">&lt;</button>
                        <button id="nextBtn" class="btn btn-primary">&gt;</button>
                    </div>
                </div>
                
                <div class="cat-showcase">
                    <div id="catCarousel" class="carousel slide" data-bs-ride="carousel">
                        <div class="carousel-inner">
                            <div class="carousel-item active">
                                <img src="/static/img/placeholder.jpg" class="d-block w-100" alt="Select a breed">
                            </div>
                        </div>
                    </div>
                    
                    <div class="breed-info mt-4">
                        <h2 id="breedName" class="text-center"></h2>
                        <p id="breedDescription" class="text-center"></p>
                        <p id="breedOrigin" class="text-center"></p>
                    </div>
                </div>
            </div>
        </div>
    </div>

    <script src="/static/js/bootstrap.bundle.min.js"></script>
    <script src="/static/js/main.js"></script>
</body>
</html>