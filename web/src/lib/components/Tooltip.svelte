<div class="tooltip">
    <slot/>
    <div class="tooltiptext">
        <slot name="tooltip"/>
    </div>
</div>

<style lang="scss">
    .tooltip {
        position: relative;
        display: inline-block;
    }

    .tooltip .tooltiptext {
        visibility: hidden;
        width: 120px;
        background-color: $grey-color-light;
        color: $light-color;
        text-align: center;
        border-radius: 6px;
        padding: 5px 0;

        position: absolute;
        z-index: 1;
    }

    .tooltip:hover .tooltiptext {
        visibility: visible;
    }
</style>