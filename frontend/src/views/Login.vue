<template>
<div>
    <out-side-nav></out-side-nav>
    <div class="container-fluid">
        <div class="row justify-content-md-center">
            <div class="col col-xs-12 col-sm-12 col-md-4 col-lg-4">
                <div class="alert alert-danger alert-dismissible fade show" role="alert">
                    <strong>Login Failed</strong>
                    <ul class="text-left">
                        <li>Invalid Username or password</li>
                    </ul>
                    <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                        <span aria-hidden="true">&times;</span>
                    </button>
                </div>
                <div class="card">
                    <div class="card-body">
                        <h4>Login</h4>
                        <div class="row">
                            <div class="col">
                                <form>
                                    <div class="form-group text-left">
                                        <label for="emailField">Email Address</label>
                                        <input type="email" class="form-control" id="emailField">
                                    </div>
                                    <div class="form-group text-left">
                                        <label for="passwordField">Password</label>
                                        <input type="password" class="form-control" id="passwordField">
                                    </div>
                                    <div class="row">
                                        <div class="col text-left">
                                            <custom-button display-text="Login" bootstrap-class="btn btn-primary"></custom-button>
                                        </div>
                                        <div class="col text-right my-auto">
                                            <a href="#">forgot password?</a>
                                        </div>
                                    </div>
                                </form>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
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
import CustomButton from '../components/buttons/CustomButton';
import OutSideNav from '../components/nav/OutSideNav';
export default {
    name:"login",
    components: {
        CustomButton,
        OutSideNav
    }
}
</script>

<style lang="scss">
    
</style>