<template>
<div>
    <nav class="navbar navbar-expand-lg navbar-dark bg-dark p-md-0">
        <div class="navbar-brand logo-area">
            <a class="" href="#">NIGHT<span>LIFE</span></a>      
        </div>
        <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
                    <span class="navbar-toggler-icon"></span>
                </button>
            
             <div class="collapse navbar-collapse" id="navbarNav">
                 <div class="search-area d-flex justify-content-center flex-fill">
                     <i><font-awesome-icon :icon="['fas', 'search']"></font-awesome-icon></i>
                     <input class="form-control" type="search" placeholder="Search" aria-label="Search">
                 </div>
                 <div class="notification-section d-flex justify-content-center mr-auto">
                                    <a href="#"><i><font-awesome-icon :icon="['far', 'envelope']"></font-awesome-icon></i></a>
                                    <a href="#"><i><font-awesome-icon :icon="['far', 'bell']"></font-awesome-icon></i></a>
                </div>
                <div class="separator"></div>
                <div class=" profile-section d-flex justify-content-center mr-auto">
                                    <img src="https://i.pravatar.cc/300" height="30" width="30" class="rounded-circle"/>
                                    <p class="user-name">Jane Doe</p>
                                    <a href="#"><i><font-awesome-icon :icon="['fas', 'caret-down']"></font-awesome-icon></i></a>
                </div>
             </div>   
    </nav>
    <div class="container-fluid dashboard-container">
        
    </div>
    <footer>
            <div class="container-fluid">
                <div class="row">
                <div class="col text-right">Copyright 2020</div>
            </div>
            </div>
    </footer>
</div>
</template>
<script>

export default {
    name: "dashboard",
}
</script>
<style lang="scss" scoped>
    a {
        color: #b2b6c3 !important;
    }
    a:hover{
        color: #fff !important;
    }
    nav {
        color:#b2b6c3;
    }
    .dashboard-container {
        min-height: 100vh;
        background: rgb(139,144,163);
        background: linear-gradient(122deg, rgba(139,144,163,1) 0%, rgba(122,126,145,1) 31%, rgba(188,194,213,1) 100%);
        color: #fff;
    }
    .logo-area {
        background-color: #143080;//#01a1ce;
        padding: 0 16px 0 16px;
        font-family: 'Monoton', cursive;
        a {
            font-size: 35px;
            color: yellow !important;
            span {
                color: #f57ff7;
                text-shadow: 2px 2px 3px #e49ab9d1;
            }
        }
    }
    .separator {
        border-left: 1px solid #b1b6c2;
        border-left-width: 1px;
        border-left-style: solid;
        border-left-color: rgb(177, 182, 194);
        height: 25px;
        padding: 0 10px 0 10px;
    }
    .user-name {
        padding-left: 10px;
        margin-top: 3px;
    }
    .profile-section {
        a {
            padding-left: 10px;
            margin-top: 3px
        }
        padding: 8px 45px 0 0;
    }
    .notification-section {
        a {
            padding-left: 10px;
            padding-right: 10px;
            font-size: 20px;
        }
        padding: 0 20px 0 20px;
    }
    .search-area {
        margin-top: 6px;
        input[type="search"] {
            background: transparent;
            border: none;
            border-radius: 0;
            box-shadow: none;
            color: #fff;
            ::placeholder {
                color: #b2b6c3;
            }
            }
        i {
            margin-top: 8px;
        }
        
    }
</style>
