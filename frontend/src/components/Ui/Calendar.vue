<template>
  <div class="calendar">
    <div class="mb-5 flex items-center justify-between">
      <ArrowNarrowLeftIcon class="w-6 h-6" @click="month--"></ArrowNarrowLeftIcon>
      <h4 class="">{{ monthName }} {{ fullYear }}</h4>
      <ArrowNarrowRightIcon class="w-6 h-6" @click="month++"></ArrowNarrowRightIcon>
    </div>
    <div class="border-main-border border-l border-t overflow-hidden">
      <div
        class="day border-main-border border-b border-r text-white"
        v-for="dayOfWeek in daysOfWeek"
        v-text="dayOfWeek"
        :key="dayOfWeek"
      />
      <div
        class="day border-main-border border-b border-r"
        v-for="emptyDay in emptyDays"
        :key="`empty-${emptyDay}`"
      />
      <div v-for="day in days" :key="day">
        <a
          class="day border-main-border border-b border-r selectable"
          :class="{
            'header-bg': formatDate(year, month, day) === today,
            selected: input.includes(formatDate(year, month, day))
          }"
          v-if="isSelectable(year, month, day)"
          v-text="day"
          @click="toggleDay(year, month, day)"
        />
        <div class="day border-b border-r" v-text="day" v-else/>
      </div>
    </div>
  </div>
</template>

<script>
import {ref, computed} from 'vue'
import {useI18n} from 'vue-i18n'
import {ArrowNarrowLeftIcon, ArrowNarrowRightIcon} from '@heroicons/vue/outline'

export default {
  name: 'Calendar',
  props: {
    modelValue: String,
  },
  emits: ['update:modelValue'],
  components: {
    ArrowNarrowLeftIcon, ArrowNarrowRightIcon,
  },
  setup(props, {emit}) {
    const {locale} = useI18n()

    const [
      yearNow = new Date().getFullYear(),
      monthNow = new Date().getMonth() + 1,
      dayNow = new Date().getDate()
    ] = props.modelValue ? props.modelValue.split('-') : [];

    const input = ref(props.modelValue);
    const year = ref(yearNow);
    const month = ref(monthNow - 1);
    const day = ref(dayNow);

    const today = computed(() => {
      return formatDate(
        new Date().getFullYear(),
        new Date().getMonth(),
        new Date().getDate()
      );
    });
    const daysOfWeek = computed(() => {
      const sunday = new Date(2017, 0, 0);
      return [...Array(7)].map(() => {
        sunday.setDate(sunday.getDate() + 1);
        return sunday.toLocaleDateString(locale, {
          weekday: 'short'
        });
      });
    });
    const monthName = computed(() => {
      const name = new Date(year.value, month.value).toLocaleString(locale, {
        month: 'long'
      });
      return `${name.charAt(0).toUpperCase()}${name.slice(1)}`;
    });
    const fullYear = computed(() => {
      return new Date(year.value, month.value).getFullYear();
    });
    const days = computed(() => {
      return new Date(year.value, month.value + 1, 0).getDate();
    });
    const emptyDays = computed(() => {
      return new Date(year.value, month.value, 1).getDay();
    });

    function formatDate(year, month, day) {
      let date = new Date(year, month, day);
      const offset = date.getTimezoneOffset();
      date = new Date(date.getTime() - offset * 60 * 1000);
      return date.toISOString().split('T')[0];
    }

    function toggleDay(year, month, day) {
      input.value = formatDate(year, month, day);
      emit('update:modelValue', input.value);
    }

    function isSelectable() {
      return true;
      /*
      const in30Days = new Date();
      in30Days.setDate(in30Days.getDate() + 30);
      return (
        new Date(year, month, day) > new Date() &&
        new Date(year, month, day) < in30Days
      );
      */
    }

    return {
      year,
      month,
      day,
      input,
      today,
      daysOfWeek,
      monthName,
      fullYear,
      days,
      emptyDays,
      toggleDay,
      formatDate,
      isSelectable
    };
  }
}
</script>

<style lang="stylus" scoped>
.calendar
  @apply mx-auto
  width 309px

  .day
    @apply text-main-text no-underline text-lg float-left text-center leading-10 cursor-pointer
    width 44px
    height 44px

    &.selectable
      @apply text-main-link bg-transparent

      &:hover
        @apply bg-main-link text-main-bg

    &.selected
      @apply bg-main-link text-main-bg

</style>
