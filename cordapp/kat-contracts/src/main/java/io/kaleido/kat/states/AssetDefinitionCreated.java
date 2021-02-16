package io.kaleido.kat.states;

import io.kaleido.kat.contracts.AssetTrailContract;
import net.corda.core.contracts.BelongsToContract;
import net.corda.core.identity.AbstractParty;
import net.corda.core.identity.Party;
import org.jetbrains.annotations.NotNull;

import java.util.List;

@BelongsToContract(AssetTrailContract.class)
public class AssetDefinitionCreated implements AssetEventState {
    private final String assetDefinitionHash;
    private final Party author;

    public String getAssetDefinitionHash() {
        return assetDefinitionHash;
    }

    public AssetDefinitionCreated(String assetDefinitionHash, Party author) {
        this.assetDefinitionHash = assetDefinitionHash;
        this.author = author;
    }

    @NotNull
    @Override
    public List<AbstractParty> getParticipants() {
        return List.of(author);
    }

    @Override
    public Party getAuthor() {
        return author;
    }
}
