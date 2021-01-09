<template>
  <div>
    <div class="sp-container">
      <sp-sign-in />
      <sp-bank-balances />
      <sp-token-send />
      <!-- this line is used by starport scaffolding # 4 -->
		<sp-type-form type="judgement" :fields="['ID', 'ChargeSheetID', 'CourtID', 'Content', 'Complete', ]" module="policingnetworkcosmos" />
		<sp-type-form type="chargesheet" :fields="['ID', 'OfficerIDs', 'FirIDs', 'InvestigationIDs', 'Content', 'Complete', ]" module="policingnetworkcosmos" />
		<sp-type-form type="evidence" :fields="['ID', 'FileExt', 'InvestigationID', ]" module="policingnetworkcosmos" />
		<sp-type-form type="investigation" :fields="['ID', 'FirID', 'OfficerID', 'Content', 'Evidence', 'Complete', ]" module="policingnetworkcosmos" />
		<sp-type-form type="fir" :fields="['ID', 'CitizenID', 'Content', 'InvestigationID', ]" module="policingnetworkcosmos" />
		<sp-type-form type="profile" :fields="['Type', 'ID', 'Name', 'Role', 'FirList', ]" module="policingnetworkcosmos" />
    </div>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";

export default {
  components: { ...sp },
};
</script>
