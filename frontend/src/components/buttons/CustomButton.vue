<template>
    <button :class="BootstrapClass">
        <div class="row justify-content-md-center">
            <div class="col my-auto">{{DisplayText}}</div>
            <div class="col">
                <fingerprint-spinner
                    :animation-duration="1000"
                    :size="30"
                    :color="'#fff'"
                />
            </div>
        </div>
    </button>
</template>

<script>
import 'epic-spinners/dist/lib/epic-spinners.min.css';
import {FingerprintSpinner} from 'epic-spinners/dist/lib/epic-spinners.min.js';
export default {
    name: "custom-button",
    components: {
        FingerprintSpinner
    },
    props: ['DisplayText', 'BootstrapClass']
}
</script>

<style lang="scss" scoped>
    .label-container {
        margin-top: 2px;
    }
</style>