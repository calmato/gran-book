<template>
  <validation-observer v-slot="{ invalid }">
    <form>
      <slot :invalid="invalid" />
    </form>
  </validation-observer>
</template>
